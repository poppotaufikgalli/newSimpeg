export default defineEventHandler(async(event) => {
    const ep = event.context.params._
    const body = await readBody(event)

    const response = await $apiClient(event, ""+ep, {
        method: "POST",
        body: body,
    })
    const {data} = response
    return data;
})