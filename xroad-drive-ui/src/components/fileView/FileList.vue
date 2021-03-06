<template>
  <v-data-table
    :headers="headers"
    :items="filteredFileList"
    :search="search"
    :group-by="groupBy ? 'dateGroup' : []"
    fixed-header
    single-select
    @click:row="clickRow"
    height="calc(100vh - 96px - 58px)"
  >
    <template v-slot:no-data>No Files</template>
    <template v-slot:item.filename="{ item }">
      <v-icon class="me-4">mdi-file-outline</v-icon>
      <span>{{ item.filename }}</span>
    </template>
    <template v-slot:item.sharedDateTime="{ item }">
      <span v-if="isToday(item.sharedDateTime)">{{ item.sharedDateTime | formatHoursMins }}</span>
      <span v-else>{{ item.sharedDateTime | formatDate }}</span>
    </template>
    <template v-slot:item.createdDateTime="{ item }">
      <span v-if="isToday(item.createdDateTime)">{{ item.createdDateTime | formatHoursMins }}</span>
      <span v-else>{{ item.createdDateTime | formatDate }}</span>
    </template>
    <template v-slot:item.filesize="{ item }">
      <span>{{ item.filesize | formatBytes }}</span>
    </template>
    <template v-slot:group.header="{ group, headers }">
      <th :colspan="headers.length">
        {{ group }}
      </th>
    </template>
  </v-data-table>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import { isToday } from '@/utils';
  import { dateGroup } from '@/filters';

  export default {
    name: 'FileList',
    props: {
      selectedFile: {
        required: true
      },
      groupBy: {
        type: String
      },
      headers: {
        type: Array
      }
    },
    computed: {
      ...mapState('fileList', ['fileList']),
      ...mapState('selectedXRoadMember', ['selectedXRoadMember']),
      ...mapState('search', ['search']),
      filteredFileList: function() {
        if (this.groupBy) {
          return this.fileList.map(item => ({
            ...item, dateGroup: dateGroup(item[this.groupBy])
          }));
        } else {
          return this.fileList;
        }
      }
    },
    methods: {
      ...mapActions('selectedFile', ['updateSelectedFile']),
      ...mapActions('fileList', ['fetchFileList']),
      ...mapActions('search', ['updateSearch']),
      clickRow(item, row) {
        if (this.selectedFile && row.isSelected) {
          row.select(false);
          this.updateSelectedFile(null);
        } else {
          row.select(true);
          this.updateSelectedFile(item);
        }
      },
      pollFileList() {
        if (this.selectedXRoadMember) {
          const serviceId = this.selectedXRoadMember + ':' + 'XRoadDrive';
          this.fetchFileList(serviceId);
        } else {
          this.fetchFileList();
        }

        this.polling = setInterval(() => {
          if (this.selectedXRoadMember) {
            const serviceId = this.selectedXRoadMember + ':' + 'XRoadDrive';
            this.fetchFileList(serviceId);
          } else {
            this.fetchFileList();
          }
        }, 3000)
      },
      isToday
    },
    created() {
      this.pollFileList()
    },
    beforeDestroy() {
      clearInterval(this.polling);
      this.updateSelectedFile(null);
      this.updateSearch('');
    }
  };
</script>

<style scoped>
  /* todo: hard-coded fixes */
  >>>.v-row-group__header {
    background: unset!important;
  }
  >>>td:not(.v-data-table__mobile-row) {
    border-bottom: unset!important;
  }
</style>
