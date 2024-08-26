export const useModalUploadDoc = defineStore('modal-upload-doc', () => {
    const displayModal = ref(false)
    const modalTitle = ref('')
    const modalAction = ref(false)
    const modalFile = ref('')
    const modalFileType = ref('')
    const fileBlob = ref(null)
    const typeAction = ref(null)

    function showModal(title, typeFile, file, ta=null) {
        modalTitle.value = title
        modalFileType.value = typeFile
        modalFile.value = file
        displayModal.value = true
        typeAction.value = ta
    }

    function closeModal(){
        modalAction.value = false
        displayModal.value = false
    }

    function doAction() {
        modalAction.value = true
        displayModal.value = false
    }

    return { 
        displayModal, 
        modalTitle, 
        modalFile, 
        modalFileType,
        showModal, 
        closeModal, 
        doAction, 
        fileBlob,
        typeAction,
    }
})