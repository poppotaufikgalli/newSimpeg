export default defineEventHandler(async(event) => {
    const body = await readBody(event)

    console.log(body)

    const response = await $apiClient(event, "/pegawai/new", {
        method: "POST",
        body: body,
    })
    const {data} = response
    return data;
})