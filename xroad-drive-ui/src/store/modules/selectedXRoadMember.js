export default {
  namespaced: true,
  state() {
    return {
      // TODO: Add comment to discribe what is content
      selectedXRoadMember: null
    };
  },
  mutations: {
    updateSelectedXRoadMember(state, selectedXRoadMember) {
      state.selectedXRoadMember = selectedXRoadMember;
    }
  },
  actions: {
    updateSelectedXRoadMember({ commit }, selectedXRoadMember) {
      commit('updateSelectedXRoadMember', selectedXRoadMember);
    }
  }
};
