<template>
  <v-container
    fluid
    class="d-flex flex-wrap mx-2 justify-center fill-height"
    flat
    tile
  >
    <v-card
      class="mx-auto"
    >
      <v-card-title>
        {{ link.title }}
      </v-card-title>
      <v-card-text>{{ link.content }}</v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import { mapState } from 'vuex'

export default {
  head () {
    return {
      title: this.link.title,
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: this.link.title
        }
      ]
    }
  },
  computed: mapState({
    link: state => state.links.link
  }),
  async fetch ({ store, error, params }) {
    try {
      await store.dispatch('links/fetchLink', params.id)
    } catch (e) {
      error({
        statusCode: 503,
        message: 'Unable to fetch link at this time.'
      })
    }
  }
}
</script>
