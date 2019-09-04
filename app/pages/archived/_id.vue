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
      title: 'Archived'
    }
  },
  components: {
    LinkCard
  },
  computed: mapState({
    links: state => state.links.links
  }),
  async fetch ({ store, error, params }) {
    try {
      await store.dispatch('links/fetchArchivedByTag', params.id)
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch links of ${params.id} at this time: ${e}`
      })
    }
  }
}
</script>
