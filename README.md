# CSVRW

CSVRW is a simple program to read CSV file output contents human-readable with nicely aligned columns.

#### Example:

'''
test ~/go/src/github.com/test/csvrw [main] $ go run main.go 

        ________  ________  ___      ___ ________  ___       __      
        |\   ____\|\   ____\|\  \    /  /|\   __  \|\  \     |\  \    
        \ \  \___|\ \  \___|\ \  \  /  / | \  \|\  \ \  \    \ \  \   
         \ \  \    \ \_____  \ \  \/  / / \ \   _  _\ \  \  __\ \  \  
          \ \  \____\|____|\  \ \    / /   \ \  \\  \\ \  \|\__\_\  \ 
           \ \_______\____\_\  \ \__/ /     \ \__\\ _\\ \____________\
                \|_______|\_________\|__|/       \|__|\|__|\|____________|
                                 \|_________|                                     
'''
'''
                                                                                                                                  
Please enter the location from the CSV file: data/sample.csv

FIRST_NAME  LAST_NAME   COMPANY_NAME                    ADDRESS                     CITY               COUNTY                STATE  ZIP    PHONE         EMAIL

James       Butt        Benton, John B Jr               6649 N Blue Gum St          New Orleans        Orleans               LA     70116  504-845-1427  jbutt@gmail.com
Josephine   Darakjy     Chanay, Jeffrey A Esq           4 B Blue Ridge Blvd         Brighton           Livingston            MI     48116  810-374-9840  josephine_darakjy@darakjy.org
Art         Venere      Chemel, James L Cpa             8 W Cerritos Ave #54        Bridgeport         Gloucester            NJ     8014   856-264-4130  art@venere.org
Lenna       Paprocki    Feltz Printing Service          639 Main St                 Anchorage          Anchorage             AK     99501  907-921-2010  lpaprocki@hotmail.com
Donette     Foller      Printing Dimensions             34 Center St                Hamilton           Butler                OH     45011  513-549-4561  donette.foller@cox.net
Simona      Morasca     Chapman, Ross E Esq             3 Mcauley Dr                Ashland            Ashland               OH     44805  419-800-6759  simona@morasca.com
Mitsue      Tollner     Morlong Associates              7 Eads St                   Chicago            Cook                  IL     60632  773-924-8565  mitsue_tollner@yahoo.com
Leota       Dilliard    Commercial Press                7 W Jackson Blvd            San Jose           Santa Clara           CA     95111  408-813-1105  leota@hotmail.com
Sage        Wieser      Truhlar And Truhlar Attys       5 Boston Ave #88            Sioux Falls        Minnehaha             SD     57105  605-794-4895  sage_wieser@cox.net
Kris        Marrier     King, Christopher A Esq         228 Runamuck Pl #2808       Baltimore          Baltimore City        MD     21224  410-804-4694  kris@gmail.com
Minna       Amigon      Dorl, James J Esq               2371 Jerrold Ave            Kulpsville         Montgomery            PA     19443  215-422-8694  minna_amigon@yahoo.com
Abel        Maclead     Rangoni Of Florence             37275 St  Rt 17m M          Middle Island      Suffolk               NY     11953  631-677-3675  amaclead@gmail.com
Kiley       Caldarera   Feiner Bros                     25 E 75th St #69            Los Angeles        Los Angeles           CA     90034  310-254-3084  kiley.caldarera@aol.com
Graciela    Ruta        Buckley Miller & Wright         98 Connecticut Ave Nw       Chagrin Falls      Geauga                OH     44023  440-579-7763  gruta@cox.net
Cammy       Albares     Rousseaux, Michael Esq          56 E Morehead St            Laredo             Webb                  TX     78045  956-841-7216  calbares@gmail.com
Mattie      Poquette    Century Communications          73 State Road 434 E         Phoenix            Maricopa              AZ     85013  602-953-6360  mattie@aol.com
Meaghan     Garufi      Bolton, Wilbur Esq              69734 E Carrillo St         Mc Minnville       Warren                TN     37110  931-235-7959  meaghan@hotmail.com
Gladys      Rim         T M Byxbee Company Pc           322 New Horizon Blvd        Milwaukee          Milwaukee             WI     53207  414-377-2880  gladys.rim@rim.org
Yuki        Whobrey     Farmers Insurance Group         1 State Route 27            Taylor             Wayne                 MI     48180  313-341-4470  yuki_whobrey@aol.com
Fletcher    Flosi       Post Box Services Plus          394 Manchester Blvd         Rockford           Winnebago             IL     61109  815-426-5657  fletcher.flosi@yahoo.com
Bette       Nicka       Sport En Art                    6 S 33rd St                 Aston              Delaware              PA     19014  610-492-4643  bette_nicka@cox.net
Veronika    Inouye      C 4 Network Inc                 6 Greenleaf Ave             San Jose           Santa Clara           CA     95111  408-813-4592  vinouye@aol.com
Willard     Kolmetz     Ingalls, Donald R Esq           618 W Yakima Ave            Irving             Dallas                TX     75062  972-896-4882  willard@hotmail.com
Maryann     Royster     Franklin, Peter L Esq           74 S Westgate St            Albany             Albany                NY     12204  518-448-8982  mroyster@royster.com
Alisha      Slusarski   Wtlz Power 107 Fm               3273 State St               Middlesex          Middlesex             NJ     8846   732-635-3453  alisha@slusarski.com
Allene      Iturbide    Ledecky, David Esq              1 Central Ave               Stevens Point      Portage               WI     54481  715-530-9863  allene_iturbide@cox.net
Chanel      Caudy       Professional Image Inc          86 Nw 66th St #8673         Shawnee            Johnson               KS     66218  913-899-1103  chanel.caudy@caudy.org
Ezekiel     Chui        Sider, Donald C Esq             2 Cedar Ave #84             Easton             Talbot                MD     21601  410-235-8738  ezekiel@chui.com
Willow      Kusko       U Pull It                       90991 Thorburn Ave          New York           New York              NY     10011  212-934-5167  wkusko@yahoo.com
Bernardo    Figeroa     Clark, Richard Cpa              386 9th Ave N               Conroe             Montgomery            TX     77301  936-597-3614  bfigeroa@aol.com
Ammie       Corrio      Moskowitz, Barry S              74874 Atlantic Ave          Columbus           Franklin              OH     43215  614-648-3265  ammie@corrio.com
Francine    Vocelka     Cascade Realty Advisors Inc     366 South Dr                Las Cruces         Dona Ana              NM     88011  505-335-5293  francine_vocelka@vocelka.com
Ernie       Stenseth    Knwz Newsradio                  45 E Liberty St             Ridgefield Park    Bergen                NJ     7660   201-387-9093  ernie_stenseth@aol.com
Albina      Glick       Giampetro, Anthony D            4 Ralph Ct                  Dunellen           Middlesex             NJ     8812   732-782-6701  albina@glick.com
Alishia     Sergi       Milford Enterprises Inc         2742 Distribution Way       New York           New York              NY     10025  212-753-2740  asergi@gmail.com
Solange     Shinko      Mosocco, Ronald A               426 Wolf St                 Metairie           Jefferson             LA     70002  504-265-8174  solange@shinko.com
Jose        Stockham    Tri State Refueler Co           128 Bransten Rd             New York           New York              NY     10011  212-569-4233  jose@yahoo.com
Rozella     Ostrosky    Parkway Company                 17 Morena Blvd              Camarillo          Ventura               CA     93012  805-609-1531  rozella.ostrosky@ostrosky.com
Valentine   Gillian     Fbs Business Finance            775 W 17th St               San Antonio        Bexar                 TX     78204  210-300-6244  valentine_gillian@gmail.com
Kati        Rulapaugh   Eder Assocs Consltng Engrs Pc   6980 Dorsett Rd             Abilene            Dickinson             KS     67410  785-219-7724  kati.rulapaugh@hotmail.com
Youlanda    Schemmer    Tri M Tool Inc                  2881 Lewis Rd               Prineville         Crook                 OR     97754  541-993-2611  youlanda@aol.com
Dyan        Oldroyd     International Eyelets Inc       7219 Woodfield Rd           Overland Park      Johnson               KS     66204  913-645-8918  doldroyd@aol.com
Roxane      Campain     Rapid Trading Intl              1048 Main St                Fairbanks          Fairbanks North Star  AK     99708  907-335-6568  roxane@hotmail.com
Lavera      Perin       Abc Enterprises Inc             678 3rd Ave                 Miami              Miami-Dade            FL     33196  305-995-2078  lperin@perin.org
Erick       Ferencz     Cindy Turner Associates         20 S Babcock St             Fairbanks          Fairbanks North Star  AK     99712  907-227-6777  erick.ferencz@aol.com
Fatima      Saylors     Stanton, James D Esq            2 Lighthouse Ave            Hopkins            Hennepin              MN     55343  952-479-2375  fsaylors@saylors.org
Jina        Briddick    Grace Pastries Inc              38938 Park Blvd             Boston             Suffolk               MA     2128   617-997-5771  jina_briddick@briddick.com
Kanisha     Waycott     Schroer, Gene E Esq             5 Tomahawk Dr               Los Angeles        Los Angeles           CA     90006  323-315-7314  kanisha_waycott@yahoo.com
Emerson     Bowley      Knights Inn                     762 S Main St               Madison            Dane                  WI     53711  608-658-7940  emerson.bowley@bowley.org
Blair       Malet       Bollinger Mach Shp & Shipyard   209 Decker Dr               Philadelphia       Philadelphia          PA     19132  215-794-4519  bmalet@yahoo.com
Brock       Bolognia    Orinda News                     4486 W O St #1              New York           New York              NY     10003  212-617-5063  bbolognia@yahoo.com
Lorrie      Nestle      Ballard Spahr Andrews           39 S 7th St                 Tullahoma          Coffee                TN     37388  931-303-6041  lnestle@hotmail.com
Sabra       Uyetake     Lowy Limousine Service          98839 Hawthorne Blvd #6101  Columbia           Richland              SC     29201  803-681-3678  sabra@uyetake.org
'''
