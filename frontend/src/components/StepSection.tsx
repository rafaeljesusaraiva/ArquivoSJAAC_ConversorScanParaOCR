import React from "react";

interface StepSectionProps {
    title: string;
    children: React.ReactNode;
    minimal?: boolean;
    end?: boolean;
}

export const StepSection: React.FC<StepSectionProps> = ({title, children, minimal, end}) => {
    var mainClass = minimal ? "col-span-4" : "col-span-full"

    return (
        <div className={mainClass}>
            <div className={"grid grid-cols-5"}>
                <div className={"col-span-full"}>
                    <h2 className={"scroll-m-20 border-b pb-2 text-xl tracking-tight first:mt-0"}>
                        {title}
                    </h2>
                </div>
                <div className={"col-span-full m-4"}>
                    {end ? children : (
                        minimal ? children : (
                            <div className={"grid grid-cols-12 items-center gap-2"}>
                                {children}
                            </div>
                        )
                    )}
                </div>
            </div>
        </div>
    )
}
