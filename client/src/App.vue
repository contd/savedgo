<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
      clipped
    >
      <h1 class="display1 ml-2 mt-2">Tags</h1>
      <div class="d-flex flex-wrap mx-2 justify-center">        
        <button
          v-for="tag in tags"
          :key="tag.label"
          class="mx-2 my-2 badge"
          :class="`badge-${tag.label}`"
          @click="filterByTag(tag.label)"
        >
          {{ tag.label }}
        </button>
      </div>

    </v-navigation-drawer>

    <v-app-bar
      app
      clipped-left
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Saved</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        v-for="navitem in navitems"
        :key="`${navitem.label}-header`"
        text
        rounded
        :to="navitem.url"
      >
        <span class="mr-2"> {{ navitem.label }} </span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>

    <v-footer app>
      {{ new Date().getFullYear() }} — <strong>Jason Kumpf</strong>
      <v-spacer></v-spacer>
    </v-footer>
  </v-app>
</template>

<script>
import '@/assets/badges.css'

export default {
  name: 'App',
  data: () => ({
    drawer: null,
    navitems: [
      {
        label: 'Home',
        url: '/'
      },
      {
        label: 'About',
        url: '/about'
      },
      {
        label: 'Login',
        url: '/login'
      },
    ],
    links: [],
    tags: [],
  }),
  methods: {
    filterByTag(tag) {
      this.$axios.get(`${this.API_BASE}/tag/${tag}`)
        .then(resp => {
          this.links = resp.data
        })
        .catch(err => {
          console.log(`AXIOS ERROR (by tag): ${err}`)
        })
    }
  },
  created() {
    this.tags = [
      {
        "label": "css",
        "topic": "developer"
      },
      {
        "label": "devops",
        "topic": "developer"
      },
      {
        "label": "dweb",
        "topic": "developer"
      },
      {
        "label": "github",
        "topic": "developer"
      },
      {
        "label": "golang",
        "topic": "developer"
      },
      {
        "label": "programming",
        "topic": "developer"
      },
      {
        "label": "javascript",
        "topic": "developer"
      },
      {
        "label": "nodejs",
        "topic": "developer"
      },
      {
        "label": "typescript",
        "topic": "developer"
      },
      {
        "label": "vuejs",
        "topic": "developer"
      },
      {
        "label": "media",
        "topic":"developer"
      },
      {
        "label": "brain",
        "topic": "humanity"
      },
      {
        "label": "health",
        "topic": "humanity"
      },
      {
        "label": "security",
        "topic": "humanity"
      },
      {
        "label": "social",
        "topic": "humanity"
      },
      {
        "label": "politics",
        "topic": "humanity"
      },
      {
        "label": "history",
        "topic": "humanity"
      },
      {
        "label": "travel",
        "topic": "humanity"
      },
      {
        "label": "awesome",
        "topic": "technology"
      },
      {
        "label": "blockchain",
        "topic": "technology"
      },
      {
        "label": "data",
        "topic": "technology"
      },
      {
        "label": "embedded",
        "topic": "technology"
      },
      {
        "label": "math",
        "topic": "technology"
      },
      {
        "label": "machinelearn",
        "topic": "technology"
      },
      {
        "label": "maker",
        "topic": "technology"
      },
      {
        "label": "linux",
        "topic": "technology"
      },
      {
        "label": "science",
        "topic": "technology"
      },
      {
        "label": "shell",
        "topic": "technology"
      },
      {
        "label": "entrepreneur",
        "topic": "utility"
      },
      {
        "label": "read",
        "topic": "utility"
      },
      {
        "label": "sites",
        "topic": "utility"
      }
    ]
  },
};
</script>
