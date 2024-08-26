export const useModalStore = defineStore('modal-store', () => {
    const displayModal = ref(false)
    const modalTitle = ref('')
    const modalText = ref('')
    const modalAction = ref(false)

    function showModal(title, text) {
        modalTitle.value = title
        modalText.value = text
        displayModal.value = true
    }

    function closeModal(){
        modalAction.value = false
        displayModal.value = false
    }

    function doAction(keyVal) {
        modalAction.value = true
        displayModal.value = false
    }

    return { 
        displayModal, 
        modalTitle, 
        modalText, 
        showModal, 
        closeModal, 
        doAction, 
    }
})