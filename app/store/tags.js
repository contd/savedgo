import tags from '@/services/tagsService'

export const state = () => ({
  tags: []
})

export const mutations = {
  SET_TAGS (state, _tags) {
    state.tags = _tags
  }
}

export const actions = {
  fetchTags ({ commit }) {
    return commit('SET_TAGS', tags)
  }
}
