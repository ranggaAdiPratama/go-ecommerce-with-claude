import { useEffect } from "react"
import { useSettings } from "./useSettings"

export const usePageTitle = (title?: string) => {
    const { settings } = useSettings()

    useEffect(
        () => {
            if (settings?.name) {
                if (title) {
                    document.title = `${settings.name} | ${title}`
                } else {
                    document.title = settings.name
                }
            }
        }, [settings, title]
    )
}