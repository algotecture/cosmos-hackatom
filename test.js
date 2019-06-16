import test from 'ava'
import execa from 'execa'

// waits number of seconds.. because cosmos is not fast enough
const wait = seconds =>
  new Promise(resolve =>
    setTimeout(resolve, seconds * 1000))

test('Alice have keys', async t => {
  const getAliceId = `nscli keys show alice -a`
  const { stdout: aliceId } = await execa.shell(getAliceId)
  // await wait(2)
  t.truthy(aliceId.startsWith('cosmos'))
})

test('Jack have keys', async t => {
  await wait(2)

  const getJackId = `nscli keys show jack -a`
  const { stdout: jackId } = await execa.shell(getJackId)
  t.truthy(jackId.startsWith('cosmos'))
})

const getAccountUrl = `http://localhost:1317/auth/accounts/`

