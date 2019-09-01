<template>
  <v-app dark>
    <!-- NAV drawer -->
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <h1 class="display1 ml-2 mt-2">
        Tags
      </h1>
      <div class="d-flex flex-wrap mx-2 justify-center">
        <button
          v-for="(tag, i) in tags"
          :key="i"
          class="mx-2 my-2 badge"
          :class="`badge-${tag.label}`"
        >
          <nuxt-link :to="`/tag/${tag.label}`">
            {{ tag.label }}
          </nuxt-link>
        </button>
      </div>
    </v-navigation-drawer>
    <!-- APP BAR -->
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer />
    </v-app-bar>
    <!-- CONTENT -->
    <v-content>
      <nuxt />
    </v-content>
    <!-- FOOTER -->
    <v-footer
      :fixed="fixed"
      app
    >
      <span>&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  head () {
    return {
      titleTemplate: '%s - Saved',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'Saved links and read later storage and annotation.'
        }
      ]
    }
  },
  data () {
    return {
      drawer: false,
      clipped: false,
      fixed: false,
      miniVariant: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Welcome',
          to: '/'
        }
      ],
      title: 'Saved',
      tags: [
        {
          'label': 'css',
          'topic': 'developer'
        },
        {
          'label': 'devops',
          'topic': 'developer'
        },
        {
          'label': 'dweb',
          'topic': 'developer'
        },
        {
          'label': 'github',
          'topic': 'developer'
        },
        {
          'label': 'golang',
          'topic': 'developer'
        },
        {
          'label': 'programming',
          'topic': 'developer'
        },
        {
          'label': 'javascript',
          'topic': 'developer'
        },
        {
          'label': 'nodejs',
          'topic': 'developer'
        },
        {
          'label': 'typescript',
          'topic': 'developer'
        },
        {
          'label': 'vuejs',
          'topic': 'developer'
        },
        {
          'label': 'media',
          'topic': 'developer'
        },
        {
          'label': 'brain',
          'topic': 'humanity'
        },
        {
          'label': 'health',
          'topic': 'humanity'
        },
        {
          'label': 'security',
          'topic': 'humanity'
        },
        {
          'label': 'social',
          'topic': 'humanity'
        },
        {
          'label': 'politics',
          'topic': 'humanity'
        },
        {
          'label': 'history',
          'topic': 'humanity'
        },
        {
          'label': 'travel',
          'topic': 'humanity'
        },
        {
          'label': 'awesome',
          'topic': 'technology'
        },
        {
          'label': 'blockchain',
          'topic': 'technology'
        },
        {
          'label': 'data',
          'topic': 'technology'
        },
        {
          'label': 'embedded',
          'topic': 'technology'
        },
        {
          'label': 'math',
          'topic': 'technology'
        },
        {
          'label': 'machinelearn',
          'topic': 'technology'
        },
        {
          'label': 'maker',
          'topic': 'technology'
        },
        {
          'label': 'linux',
          'topic': 'technology'
        },
        {
          'label': 'science',
          'topic': 'technology'
        },
        {
          'label': 'shell',
          'topic': 'technology'
        },
        {
          'label': 'entrepreneur',
          'topic': 'utility'
        },
        {
          'label': 'read',
          'topic': 'utility'
        },
        {
          'label': 'sites',
          'topic': 'utility'
        }
      ]
    }
  }
}
</script>
