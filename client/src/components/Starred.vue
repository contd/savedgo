<template>
  <svg
    width="20"
    height="20"
    viewBox="0 0 20 20"
    xmlns="http://www.w3.org/2000/svg"
    data-svg="star"
    @click="toggle"
  >
    <polygon
      :fill="fill"
      :stroke="stroke"
      stroke-width="1.01"
      points="10 2 12.63 7.27 18.5 8.12 14.25 12.22 15.25 18 10 15.27 4.75 18 5.75 12.22 1.5 8.12 7.37 7.27">
    </polygon>
  </svg>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Starred',
  props: {
    linkId: String
  },
  data() {
    return {
      isStarred: false,
      stroke: '#999',
      fill: 'none',
      url: `/starred/${this.linkId}`
    }
  },
  methods: {
    toggle: function() {
      if (this.isStarred === false) {
        this.isStarred = true
        this.fill = '#FFDC00'
      } else {
        this.isStarred = false
        this.fill = 'none'
      }

      axios.put(`${this.$baseUrl}${this.url}`, {
        is_starred: this.isStarred
      })
        .then(resp => {
          console.log(`Toggled is_starred: ${resp.data.is_starred}`)
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
</script>

<style scoped>
svg {
  display: inline-block;
  vertical-align: baseline;
  margin-bottom: -2px;
  cursor: pointer;
}
</style>