import http from "k6/http";

export let options = {
  vus: 2,
  duration: "2s",
};

export default function () {
  //   http.get("http://localhost:8080/hello");
  http.get("http://host.docker.internal:8080/products");
}
