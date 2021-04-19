import service from '@/utils/request'

export const importCSVtoDB = (data) => {
    return service({
        url: "/dataManage/importCSVtoDB",
        method: 'put',
        data
    })
}