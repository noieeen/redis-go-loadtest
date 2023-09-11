import http from "k6/http";

export const options = {
  vus: 10,
  duration: "11s",
};

export default function () {
  http.get("http://host.docker.internal:8000/products");
}
