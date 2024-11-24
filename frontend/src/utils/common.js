export const CommonVerification = (obj) => {
    if (Object.prototype.toString.call(obj) === '[object Object]') {
        for (const objKey in obj) {
            if (obj[objKey] === "") {
                return false
            }
        }
        return true
    } else {
        console.error("CommonVerification error: obj 类型不正确")
    }
}

export const getTableTypeName = (typeId) => {
    switch (typeId) {
        case "1":
            return "气象基础设施资源池资源申请表";
        case "2":
            return "资源需求预估表";
        case "3":
            return "气象基础设施资池资源变更表";
    }
}

export const getStatusText = (status) => {
    switch (status) {
        case "-99":
            return "驳回";
        case "-1":
            return "草稿";
        case "0":
            return "提交";
        case "1":
            return "一级审批";
        case "10":
            return "办结";
    }
}

export const showActionBtn = (status) => {
    return status > 0 && status < 10
}