export default {
  namespaced: true,
  state() {
    return {
      // TODO: Add comment to discribe what is content
      selectedFile: null
    };
  },
  mutations: {
    updateSelectedFile(state, selectedFile) {
      state.selectedFile = selectedFile;
    }
  },
  actions: {
    updateSelectedFile({ commit }, selectedFile) {
      commit('updateSelectedFile', selectedFile);
    }
  }
};
