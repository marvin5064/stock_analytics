library(quantmod)

stocks<-c("0941.HK", "0700.HK", "3323.HK")
getSymbols(stocks,src="yahoo",from="2010-01-01",auto.assign = T)



plot_Stock_Since_Year <- function(stock_Name, input_Year) {
    current_Year <- as.integer(format(Sys.Date(), "%Y"))
    for (year in input_Year:current_Year) {
        # year graph
        # plot_Year_Graph(year)
        print(sprintf("year %d", year))
        # within 2 years
        if (year > current_Year - 3) {
            # month graph
            for (month in 1:12) {
                current_Month <- as.integer(format(Sys.Date(), "%m"))
                if (current_Year == year && month > current_Month)  {
                    next
                }
                # plot_Month_Graph()
                print(sprintf("year %d month %d", year, month))
            }
        }
    }
}


# chart_Series(`3323.HK`)
# zoom_Chart("2010")
# zoom_Chart("2011")
# zoom_Chart("2012")
# zoom_Chart("2013")
# zoom_Chart("2014")
# zoom_Chart("2015")
# zoom_Chart("2016")
# zoom_Chart("2017")
# zoom_Chart("2018")
# zoom_Chart("2018-01")
# zoom_Chart("2018-02")
# zoom_Chart("2018-03")

# # SMA19<-SMA(Cl(`0700.HK`),n=19)
# # add_TA(SMA19,on=1,col="blue")
# # SMA9<-SMA(Cl(`0700.HK`),n=9)
# # add_TA(SMA9,on=1,col="red")
# # SMA50<-SMA(Cl(`0700.HK`),n=50)
# # add_TA(SMA50,on=1,col="yellow")

# # RSI9<-RSI(Cl(`0700.HK`),n=9)
# # add_TA(RSI9,col="black")


# plot_Month_Graph <- function(month, year, output) {

# }

# plot_Year_Graph <- function(year, output) {

# }

