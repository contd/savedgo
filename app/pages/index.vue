<template>
  <v-container
    fluid
    class="d-flex flex-wrap mx-2 justify-center fill-height"
    flat
    tile
  >
    <LinkCard
      v-for="(link, index) in links"
      :key="index"
      :link="link"
    />
  </v-container>
</template>

<script>
import { mapState } from 'vuex'
import LinkCard from '@/components/LinkCard.vue'

export default {
  head () {
    return {
      title: 'Home'
    }
  },
  components: {
    LinkCard
  },
  computed: mapState({
    links: state => state.links.links
  }),
  async fetch ({ store, error }) {
    try {
      await store.dispatch('links/fetchLinks')
    } catch (e) {
      error({
        statusCode: 503,
        message: 'Unable to fetch starred links at this time.'
      })
    }
  }
}
</script>
