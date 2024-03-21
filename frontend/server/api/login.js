export default defineEventHandler(async(event) => {
    const body = await readBody(event)

    const config = useRuntimeConfig()
    const formdata = new FormData();
    formdata.append("username", body.username);
    formdata.append("password", body.password);
    //console.log(body)
    return fetch(config.public.apiBase+"/login", {
        method: 'POST',
        body: formdata,
    }).then((response) => 
        response.json().then(data => ({status: response.status, data: data}))
    )
})