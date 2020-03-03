import unfetch from 'unfetch'

const host = "http://localhost:8081"

export const get = (endpoint) => {
  unfetch(`${host}/${endpoint}`).then()
}