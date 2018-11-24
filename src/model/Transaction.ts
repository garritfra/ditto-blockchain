export default class Transaction {
  sender: any;
  receiver: any;
  amount: any;

  constructor(sender, receiver, amount) {
    this.sender = sender;
    this.receiver = receiver;
    this.amount = amount;
  }
}
