import fetch from 'unfetch'
import { Item } from '../proto/item_pb'

class ItemHandler {
  private _host: string

  constructor(host: string) {
    this._host = host
    this.getItemById = this.getItemById.bind(this)
  }

  async getItemById(id: number): Promise<Item.AsObject> {
    return fetch(`${this._host}/v1/item/${id}`).then(res => res.json())
  }
}

export default ItemHandler