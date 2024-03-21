export default defineEventHandler(async(event) => {
    const response = await $apiClient(event, "/pegawai")
    const {data} = response
    return data;
})