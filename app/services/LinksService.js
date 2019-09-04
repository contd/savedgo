import axios from 'axios'

const apiClient = axios.create({
  baseURL: 'http://localhost:4444',
  withCredentials: false,
  debug: true,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getLinks () {
    return apiClient.get('/starred')
  },
  getLink (id) {
    return apiClient.get('/id/' + id)
  },
  getLinksByTag (tag) {
    return apiClient.get('/tag/' + tag)
  },
  getArchivedByTag (tag) {
    return apiClient.get('/archived/' + tag)
  }
}
