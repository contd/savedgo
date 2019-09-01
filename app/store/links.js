import LinksService from '@/services/LinksService.js'

export const state = () => ({
  links: [],
  link: {},
  tags: []
})

export const mutations = {
  SET_LINKS (state, links) {
    state.links = links
  },
  SET_LINK (state, link) {
    state.link = link
  },
  SET_TAGS (state, tags) {
    state.tags = tags
  }
}

export const actions = {
  fetchLinks ({ commit }) {
    return LinksService.getLinks().then((resp) => {
      commit('SET_LINKS', resp.data)
    })
  },
  fetchLink ({ commit }, id) {
    return LinksService.getLink(id).then((resp) => {
      commit('SET_LINK', resp.data)
    })
  },
  fetchLinksByTag ({ commit }, tag) {
    return LinksService.getLinksByTag(tag).then((resp) => {
      commit('SET_LINKS', resp.data)
    })
  }
}
