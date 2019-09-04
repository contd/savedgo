export default ({ $axios }) => {
  $axios.onRequest((config) => {
    console.log('Axios Making request to: ' + config.url)
  })

  $axios.onRequestError((err) => {
    console.log('Axios Error Making request: ' + err)
  })

  $axios.onError((err) => {
    console.log('Axios Error: ' + err)
  })
}
