import http from "k6/http";

export const options = {
  vus: 10,
  duration: "60s",
//   iterations: 1,

  //   stages: [
  // { duration: "10s", target: 10 },
  //   { duration: '1m30s', target: 10 },
  //   { duration: '20s', target: 0 },
  //   ],
};

export default function () {
  http.get("http://host.docker.internal:8000/hello");
}
