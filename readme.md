# insurance
**insurance** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli). It allows you to add car damage records.

Versions used are:

* Go: 1.18.3
* Ignite CLI: 0.22.2
* Cosmos SDK: 0.45.4

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).


## Adding records

* To add new record use:

```
./insuranced tx insurance add-record [brand] [model] [year] [owner] [licensePlate] [vinNumber] [odometer] [side] [part] [payout] --from insurance_company
```
`insurance_company` should match your account name


## Viewing records

* To view all added records use:

```
./insuranced query insurance list-crash-rec
```
*To view added record by it's index use:

```
./insuranced query insurance show-crash-rec recNum
```
`recNum` should match index of record you want to view
