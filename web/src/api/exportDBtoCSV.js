import service from '@/utils/request'

export const exportDBtoCSV = (data) => {
    return service({
        url: "/dataManage/exportDBtoCSV",
        method: 'post',
        data
    })
}