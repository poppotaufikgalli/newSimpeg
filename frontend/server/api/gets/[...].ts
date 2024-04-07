export default defineEventHandler(async(event) => {
    const ep = event.context.params._
    const response = await $apiClient(event, "/"+ep)
    const {data} = response
    return data;
})