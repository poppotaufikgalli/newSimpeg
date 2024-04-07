export default defineEventHandler(async(event) => {
    const res = await $apiClient(event, "/islogin")
    return res;
})

/*export default defineEventHandler(async(event) => {
    const {token} = parseCookies(event)
    const config = useRuntimeConfig();

    let headers = {
        "Authorization": "Bearer " + token,
        "Accept": "application/json",
    };

    return fetch(config.public.apiBase+"/islogin", {
        headers: headers
    }).then((response) => 
        response.json().then(data => ({status: response.status, data: data}))
    )
})*/