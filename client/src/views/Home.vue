<template>
  <div>
    <v-container fluid
      class="d-flex flex-wrap mx-2 justify-center fill-height"
      flat
      tile
    >
      <div v-for="link in links"
          :key="link._id"
      >
        <v-hover v-slot:default="{ hover }">
          <v-card           
            class="mx-1 my-1"
            width="400"
            :elevation="hover ? 12 : 2"
            :class="{ 'on-hover': hover }"
          >
            <v-img
              class="white--text"
              height="200px"
              :src="link.preview_picture"
            >          
            </v-img>

            <v-card-title class="title">
              <v-tooltip bottom>
                <template v-slot:activator="{ on }">
                  <span v-on="on" class="truncate">{{ link.title }}</span>
                </template>
                <span>{{ link.title }}</span>
              </v-tooltip>
            </v-card-title>
            
            <v-card-text>
              <span class="badge" :class="badgeClass(link.tags[0])">
                {{ link.tags[0] }}
              </span>
            </v-card-text>

            <v-divider></v-divider>
            
            <v-card-actions
              color="cyan"
              dark
            >
              <v-btn 
                text
                target="_new"
              >
                <v-icon>mdi-link</v-icon>
                {{ link.domain_name }}
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn icon>
                <v-icon>mdi-star</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-bookmark</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-delete</v-icon>
              </v-btn>
            </v-card-actions>

            <v-expand-transition>
              <div v-show="show">
                <v-card-text> 
                </v-card-text>
              </div>
            </v-expand-transition>
          </v-card>
        </v-hover>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data: () => ({
    links: []
  }),
  methods: {
    badgeClass(tag) {
      return `badge-${tag}`
    }
  },
  created() {
    this.$axios.get(`${this.$API_BASE}/starred`)
      .then(resp => {
        this.links = resp.data
      })
      .catch(err => {
        console.log(`AXIOS ERROR (starred): ${err}`)
      })
  }
};
</script>

<style lang="css">
.truncate {
  width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>