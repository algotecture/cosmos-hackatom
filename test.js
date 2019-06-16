import test from 'ava'
import execa from 'execa'
import axios from 'axios'
import dagOfAlexPlatz from './data/dag.json'

let aliceId
let jackId

const getLocationsUrl = streetAndNumber =>
  streetAndNumber
    ? `http://localhost:1317/nameservice/locations/${streetAndNumber}/whois`
    : `http://localhost:1317/nameservice/locations`

// waits number of seconds.. because cosmos is not fast enough
const wait = seconds =>
  new Promise(resolve =>
    setTimeout(resolve, seconds * 1000))

test('Alice have keys', async t => {
  const getAliceId = `nscli keys show alice -a`
  const resp = await execa.shell(getAliceId)
  aliceId = resp.stdout
  // await wait(2)
  t.truthy(aliceId.startsWith('cosmos'))
})

test('Jack have keys', async t => {
  await wait(2)

  const getJackId = `nscli keys show jack -a`
  const resp = await execa.shell(getJackId)
  jackId = resp.stdout

  t.truthy(jackId.startsWith('cosmos'))
})


test('Alexander platz is taken', async t => {
  const { data: alexPlatz } = await axios.get(getLocationsUrl('AlexPlatz1'))

  await wait(3)
  const getJackId = `nscli keys show jack -a`
  const resp = await execa.shell(getJackId)
  jackId = resp.stdout

  t.truthy(alexPlatz.owner)
  t.deepEqual(alexPlatz.owner, jackId, 'Alex')


  t.deepEqual(JSON.parse(alexPlatz.value), dagOfAlexPlatz)
  // TODO lat lon not empty
})

test('Get all locations', async t => {
  const { data: locations } = await axios.get(getLocationsUrl())

  t.deepEqual(locations, ['AlexPlatz1', 'BrandenburgerTor1'])
})

// TODO show buildings/ locations
// $ curl -s http://localhost:1317/nameservice/names/jack1.id
const getAccountUrl = `http://localhost:1317/auth/accounts/`
