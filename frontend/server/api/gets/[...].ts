export default defineEventHandler(async(event) => {
    let ep = event.context.params._
    const {searchNama, limit} = getQuery(event)

    if(searchNama != undefined){
        ep = ep + "?searchNama="+searchNama

        if(limit != undefined){
            ep = ep + "&limit="+limit
        }
    }

    const response = await $apiClient(event, "/"+ep)
    const {data} = response
    return data;
})