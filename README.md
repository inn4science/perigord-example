**This project is an example of using Perigord for smart-contract testing**

The contract Market is an example of an owned contract that cann accept ETH from predefined wallets and return this ETH to a designated person upon the call of the contract owner.

##Running tests

Please, follow [Enhanced Perigord](https://gitlab.com/go-truffle/enhanced-perigord) installation and running instructions.

Clone this project to src/ folder in your GOPATH.

Start Ganache on HTTP://127.0.0.1:7545

Go to the project directory and run:

```$xslt
$ perigord test
```