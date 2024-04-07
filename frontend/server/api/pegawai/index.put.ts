export default defineEventHandler(async(event) => {
    const body = await readBody(event)

    const response = await $apiClient(event, "/pegawai", {
        method: "PUT",
        body: body,
    })
    const {data} = response
    return data;
})