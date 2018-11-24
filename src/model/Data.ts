import { RequestType } from "./RequestType";

export default class Data {
  requestType: RequestType;
  data: any;

  constructor(requestType: RequestType, data: any) {
    this.requestType = requestType;
    this.data = data;
  }
}
