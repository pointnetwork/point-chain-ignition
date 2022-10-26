import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateDelegatedAmount } from "./types/point/lockenomics/tx";
import { MsgUpdateDelegatedAmount } from "./types/point/lockenomics/tx";
import { MsgDeleteDelegatedAmount } from "./types/point/lockenomics/tx";
import { MsgCreateLock } from "./types/point/lockenomics/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/point.lockenomics.MsgCreateDelegatedAmount", MsgCreateDelegatedAmount],
    ["/point.lockenomics.MsgUpdateDelegatedAmount", MsgUpdateDelegatedAmount],
    ["/point.lockenomics.MsgDeleteDelegatedAmount", MsgDeleteDelegatedAmount],
    ["/point.lockenomics.MsgCreateLock", MsgCreateLock],
    
];

export { msgTypes }