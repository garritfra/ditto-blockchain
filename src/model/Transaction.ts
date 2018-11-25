export default class Transaction {
  sender: string;
  receiver: string;
  amount: number;

  constructor(sender: string, receiver: string, amount: number) {
    this.sender = sender;
    this.receiver = receiver;
    this.amount = amount;
  }
}
