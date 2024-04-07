export default defineEventHandler(async(event) => {
    const response = await $apiClient(event, "/pegawai/list")
    const {data} = response
    return data;
})