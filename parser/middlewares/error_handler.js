
export const errorHandler = async (err, req, res, next) => {
    const status = err.status
    if (status !== undefined) {
        delete err.status
        res.status(status).json(err )
        return
    }
    console.log(err);
    res.status(500).json({error:"Something went wrong"})
}