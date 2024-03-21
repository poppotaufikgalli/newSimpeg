export default function (event, request, opts) {
    const config = useRuntimeConfig()

    const {token} = parseCookies(event)

    let headers = {
        "Authorization": "Bearer " + token,
        "Accept": "application/json",
        ...opts?.headers,
    };

    return $fetch(request, {
        baseURL: config.public.apiBase,
        ...opts,
        headers: headers
    })
}