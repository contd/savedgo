const BASE_URL = "http://localhost:4444";

Vue.component("card-item", {
  template: "#card-template",
  props: {
    title: String,
    created_at: String,
    reading_time: String,
    tags: Array,
    domain_name: String,
    is_starred: Boolean,
    is_archived: Boolean
  }
});

const app = new Vue({
  el: "#app",
  data() {
    return {
      links: [],
      archive: false
    };
  },
  mounted() {
    axios
      .get(`${BASE_URL}/starred`)
      .then(resp => {
        this.links = resp.data;
      })
      .catch(err => {
        console.log(`Axios Error: ${err}`);
      });
  },
  computed: {
    badgeHref(label) {
      if (this.archive) {
        return `/archive/tag/${label}`;
      } else {
        return `/tag/${label}`;
      }
    },
    badgeClass(label) {
      return `badge badge-${label}`;
    }
  }
});
