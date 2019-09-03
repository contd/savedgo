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
      <v-card-title class="display-1">
        {{ link.title }}
      </v-card-title>

      <v-card-text>
        <div class="d-flex blue lighten-5 px-3 py-3">
          <span class="grey--text subtitle-1 mx-2 pt-1">
            <v-icon>mdi-calendar</v-icon>
            {{ link.created_at | humanDate }}
          </span>
          <span class="grey--text subtitle-1 mx-2 pt-1">
            <v-icon>mdi-clock</v-icon>
            {{ link.reading_time }} min
          </span>
          <div class="flex-grow-1" />
          <button
            v-for="(tag, idx) in link.tags"
            :key="idx"
            class="mx-2 badge"
            :class="`badge-${tag}`"
          >
            {{ tag }}
          </button>
          <v-btn
            class="mx-2"
            @click="goBack"
          >
            <v-icon>mdi-arrow-left</v-icon>
            Back
          </v-btn>
          <v-btn
            class="mx-2"
            :href="link.url"
            target="_new"
          >
            <v-icon>mdi-link</v-icon>
            Original Link
          </v-btn>
        </div>
        <div
          class="article"
          v-html="link.content"
        />
      </v-card-text>

      <v-card-actions>
        <v-btn
          text
          @click="goBack"
        >
          Back
        </v-btn>
        <v-btn
          text
          :href="link.url"
          target="_new"
        >
          Original Link
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import { mapState } from 'vuex'
import moment from 'moment'

export default {
  middleware: 'auth',
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
  filters: {
    humanDate (date) {
      return moment(date).format('ddd, MMM Do YYYY')
    },
    codeReplace (content) {
      let replContent = content.replace(/<pre class="(\w+)">/g, '<pre v-highlightjs><code class="$1">')
      replContent = replContent.replace(/<\/pre>/g, '</code></pre>')
      return replContent
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
  },
  methods: {
    goBack () {
      this.$router.go(-1)
    }
  }
}
</script>
