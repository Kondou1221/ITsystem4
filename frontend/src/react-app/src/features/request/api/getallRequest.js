
const config = {
    apiPrefix: process.env.NEXT_PUBLIC_API_PREFIX ?? "http://",
    apiHost: process.env.NEXT_PUBLIC_VERCEL_URL ?? "localhost:3000"
}

export default async function Searchrequest() {
    const url = config.apiPrefix + config.apiHost + "/request_get"
    // const url = "http://localhost:3030/request_get"
    // console.log(url)
    const response = await fetch("http://back_go:8080/request_get")
    const res = response.json()
    // console.log(res)
    return res
}
