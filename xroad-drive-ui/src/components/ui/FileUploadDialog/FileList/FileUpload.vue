<script>
  import { mapState, mapActions } from "vuex";

  export default {
    name: 'FileUpload',
    props: {
      file: {
        type: Object,
        required: true
      }
    },
    computed: {
      ...mapState('selectedXRoadMember', ['selectedXRoadMember'])
    },
    methods: {
      ...mapActions('uploadFiles', [
        'upload',
        'updateProgressById',
        'updateIsDoneById',
        'updateIndeterminateById'
      ])
    },
    async mounted() {
      const { file } = this;
      const id = file.id;

      let serviceId;
      if (this.selectedXRoadMember) {
        serviceId = this.selectedXRoadMember + ':' + 'XRoadDrive';
      }

      if (!file) return;
      console.log(`Start to upload file: ${file.name}`)
      try {
        await this.upload({uploadFile: file, serviceId});
        this.updateIsDoneById({ id, isDone: true });
        this.updateProgressById({ id, progress: 100 });
        console.log(`Succeeded to upload file: ${file.name}`);
      } catch (error) {
        this.updateIsDoneById({ id, isDone: false });
        this.updateProgressById({ id, progress: 100 });
        this.updateIndeterminateById({ id, indeterminate: false });
        console.log(`Failed to upload file: ${file.name}`);
        console.log(error);
      }
    },
    render: () => null,
  };
</script>
