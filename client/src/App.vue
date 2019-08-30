<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-toolbar-title class="headline text-uppercase">
        <span>Saved</span>
        <span class="font-weight-light"></span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        v-for="link in links"
        :key="`${link}-header`"
        text
        rounded
        :to="link.url"
      >
        <span class="mr-2"> {{ link.label }} </span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>

    <v-footer
      color="primary lighten-1"
      padless
    >
      <v-row
        justify="center"
        no-gutters
      >
        <v-btn
          v-for="link in links"
          :key="`${link}-footer`"
          color="white"
          text
          rounded
          class="my-2"
          :to="link.url"
        >
          {{ link.label }}
        </v-btn>
        <v-col
          class="primary lighten-2 py-4 text-center white--text"
          cols="12"
        >
          {{ new Date().getFullYear() }} — <strong>Jason Kumpf</strong>
        </v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    links: [
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
    ]
  }),
};
</script>
