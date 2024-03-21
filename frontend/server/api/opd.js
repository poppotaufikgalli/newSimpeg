export default defineEventHandler(async(event) => {
    const response = await $apiClient(event, "/opd")
    const {data} = response
    return data;
})