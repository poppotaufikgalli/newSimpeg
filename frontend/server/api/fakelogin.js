export default defineEventHandler(async(event) => {
    const body = await readBody(event)

    const config = useRuntimeConfig()
    return fetch(config.public.apiBase+"/fakelogin", {
        method: 'POST',
    }).then((response) => 
        response.json().then(data => ({status: response.status, data: data}))
    )
})