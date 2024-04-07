export default defineEventHandler(async(event) => {
    const body = await readBody(event)

    const response = await $apiClient(event, "/pegawai/list", {
        method: "POST",
        body: body,
    })
    const {data} = response
    return data;
})