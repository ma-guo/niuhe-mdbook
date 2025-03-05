
interface DataSet {
    [key: string]: any
}

interface TapDetail {
    x: number
    y: number
    /**  用作输入框时使用 */
    value: any
    source: string
    /** change 事件 */
    current: number
    // userInfo?: WechatMiniprogram.UserInfo
    causedBy: string
    rotate: number
    scale: number
    skew: number
    type: string
    markerId: number
}

interface TapTarget {
    id: string
    offsetLeft: number
    offsetTop: number
    dataset: DataSet
}

interface TapTouche {
    identifier: number
    pageX: number
    pageY: number
    clientX: number
    clientY: number
    force: number
}

/** 在 wx 定义中找不到 tap 事件传进来的参数类型, 为了可读, 在这里将 log 出来的结构信息定义如下 */
interface TapEvent {
    type: string
    timeStamp: number
    target: TapTarget
    currentTarget: TapTarget
    detail: TapDetail
    touches: TapTouche[]
    changedTouches: TapTouche[]
    mut: boolean
    _userTap: boolean
}

interface TriggerEvent {
    type: string
    timeStamp: number
    target: TapTarget
    currentTarget: TapTarget
    detail: any
    touches: TapTouche[]
    changedTouches: TapTouche[]
    mut: boolean
}
/** 通用返回结构 */
interface Rsp<T> {
    result: number
    errMsg: string
    data: T
    // header?: WechatMiniprogram.IAnyObject
}

