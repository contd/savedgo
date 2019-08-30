<template>
  <span>
    <span v-for="tag in tags"
      :key="tag"
      class="badge"
      :class="badgeClass(tag)"
    >
      {{tag}}
    </span>
  </span>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Tags',
  props: {
    linkId: String
  },
  data() {
    return {
      tags: []
    }
  },
  computed: {
    badgeClass(tag) {
      return `badge-${tag}`
    }
  },
  mounted() {
    axios.get(`${this.$baseUrl}/${this.linkId}`)
      .then(resp => {
        this.tags = resp.data.tags
      })
      .catch(err => {
        this.tags = []
        console.log(`Error getting link tags: ${err.message}`)
      })
  }
}
</script>