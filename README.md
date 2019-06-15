# SDK Application Tutorial

This repository contains the source code of the nameservice tutorial.

## Tutorial

**[Click here](./tutorial/README.md)** to access the tutorial. You can also view it on the [website](https://cosmos.network/docs/tutorial).

## Building and running the example

**[Click here](./tutorial/build-run.md)**  for instructions on how to build and run the code.

Translations:
- [中文](./README_cn.md)

## [Slides](https://docs.google.com/presentation/d/1aCMAdkVY-gfgnGNPTygwVk3o68czPQ_VYfvdMy9Ek5Q/edit?usp=sharing)


## Bellow this line are notes
--

## go mod
```bash
go mod replace... can be by hand
go mod verify
go mod tidy
go build
```


## DAG
```
ipfs hash
QmWHTHBapomvBksnuAWSQZ47ckj3LrGgmb7D2r1WXExNEX
ipfs daemon &
ipfs cat QmWHTHBapomvBksnuAWSQZ47ckj3LrGgmb7D2r1WXExNEX
```

## TODO discuss things
- [ ] Branding internet of Buildings
- [ ] Purchasing buildings
    - [ ] split ownership as CrowdHouse
    - should it be in DAG.. maybe in the chain ?
    - [X] this in DAG `OWNERHISTORY/ApplicationIdentifier`
        - should go to to ownership chain
    - is it enough to remove notarial office
- [ ] TDD approach to buildings with smart contracts
    - automatically recalculate stress tests,
    - automatically recreate energy norms
    - automatically recreate ISO-Norms
- [ ] Forking buildings, and paying for DAG to clone
    - [ ] should DAG be private
        - maybe a private chain for it ?
    - [ ] how enforce payments of for DAG's 
    - [X] how to encrypt DAG's
        - DAG can be private 
    - [X] is IPFS good place to handle DAG?   
        - maybe just use private chain
