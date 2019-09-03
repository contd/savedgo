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
      <v-list disabled>
        <v-subheader class="display-1">
          Tags
        </v-subheader>
        <v-list-item-group color="primary">
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
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
    <!-- APP BAR -->
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title class="display-1" v-text="title" />
      <div class="flex-grow-1" />
      <v-row justify="center">
        <v-dialog v-model="dialog" persistent max-width="600px">
          <template v-slot:activator="{ on }">
            <v-btn color="primary" dark v-on="on">
              Add Link
            </v-btn>
          </template>
          <v-card>
            <v-card-title class="headline">
              Add New Link
            </v-card-title>
            <v-card-text>
              <v-form class="mt-2">
                <v-row>
                  <v-col>
                    <v-text-field
                      v-model="linkTitle"
                      v-validate="'required|between:3,255'"
                      v-slot="errors"
                      label="Title"
                      outlined
                      data-vv-name="title"
                      required
                      prepend-icon="mdi-textbox"
                    />
                    <span>{{ errors[0] }}</span>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col>
                    <v-text-field
                      v-model="linkUrl"
                      v-validate="'required|between:3,255'"
                      v-slot="errors"
                      label="Url"
                      outlined
                      data-vv-name="url"
                      required
                      prepend-icon="mdi-link"
                    />
                    <span>{{ errors[0] }}</span>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col>
                    <v-combobox
                      v-model="linkTag"
                      :items="tags"
                      item-text="label"
                      item-value="label"
                      :search-input.sync="search"
                      small-chips
                      multiple
                      hide-selected
                      clearable
                      prepend-icon="mdi-tag"
                      label="Tag"
                      outlined
                      data-vv-name="tag"
                    />
                  </v-col>
                </v-row>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <div class="flex-grow-1" />
              <v-btn color="green darken-1" text @click="submit()">
                Submit
              </v-btn>
              <v-btn color="green darken-1" text @click="dialog = false">
                Cancel
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
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
import { extend } from 'vee-validate'
import { required, email } from 'vee-validate/dist/rules'
import tags from '@/services/tags'

extend('required', required)
extend('email', email)

export default {
  $_veeValidate: {
    validator: 'new'
  },
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
      clipped: true,
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
      tags,
      linkTitle: '',
      linkUrl: '',
      linkTag: '',
      snackbar: false,
      errors: [],
      errorMsg: '',
      search: null,
      dialog: false
    }
  },
  methods: {
    submit () {
      this.$validator.validate().then((valid) => {
        if (!valid) {
          // Do something to show errors
        } else {
          this.dialog = false
          // axios POST /
        }
      })
    }
  }
}
</script>
