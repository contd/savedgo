<template>
  <svg
    width="20"
    height="20"
    viewBox="0 0 20 20"
    xmlns="http://www.w3.org/2000/svg"
    data-svg="check"
    @click="toggle"
  >
    <polyline
      :fill="fill"
      :stroke="stroke"
      stroke-width="1.1"
      points="4,10 8,15 17,4"
    ></polyline>
  </svg>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Archived',
  props: {
    isArchived: Boolean,
    linkId: String
  },
  data() {
    return {
      state: this.isArchived,
      stroke: '#999',
      fill: 'none',
      url: `/archived/${this.linkId}`
    }
  },
  methods: {
    toggle: function() {
      if (this.state === false) {
        this.state = true
        this.isArchived = true
        this.fill = '#FFDC00'
      } else {
        this.state = false
        this.isArchived = false
        this.fill = 'none'
      }

      axios.put(`${this.$baseUrl}${this.url}`, {
        is_archived: this.isArchived
      })
        .then(resp => {
          console.log(`Toggled is_archived: ${resp.data.is_archived}`)
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