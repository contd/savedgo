<template>
  <v-card
    class="mx-1 my-1"
    width="400"
  >
    <v-img
      class="white--text"
      height="200px"
      :src="link.preview_picture"
    />

    <nuxt-link :to="`/link/${link._id}`">
      <v-card-title class="title">
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <span class="truncate" v-on="on">{{ link.title }}</span>
          </template>
          <span>{{ link.title }}</span>
        </v-tooltip>
      </v-card-title>
    </nuxt-link>

    <v-card-text>
      <button
        v-for="(tag, idx) in link.tags"
        :key="idx"
        class="badge"
        :class="`badge-${tag}`"
      >
        {{ tag }}
      </button>
    </v-card-text>

    <v-divider />

    <v-card-actions
      color="cyan"
      dark
    >
      <v-btn
        text
        target="_new"
        :to="link.url"
      >
        <v-icon>mdi-link</v-icon>
        {{ link.domain_name }}
      </v-btn>

      <v-spacer />

      <v-btn icon>
        <v-icon>{{ link.is_starred ? 'mdi-star' : 'mdi-star-outline' }}</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>{{ link.is_archived ? 'mdi-bookmark' : 'mdi-bookmark-outline' }}</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <div v-show="show">
        <v-card-text />
      </div>
    </v-expand-transition>
  </v-card>
</template>

<script>
import moment from 'moment'

export default {
  name: 'LinkCard',
  filters: {
    humanDate (date) {
      return moment(date).format('ddd, MMM Do YYYY')
    }
  },
  props: {
    link: {
      type: Object,
      default: null
    }
  },
  data: () => ({
    show: false,
    tags: []
  }),
  methods: {
    badgeClass (tag) {
      return `badge-${tag}`
    }
  }
}
</script>

<style lang="css">
.truncate {
  width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
